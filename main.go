package main

import (
	"fmt"
	"math"
)

// 圆结构体
type Circle struct {
	Radius float64 // 半径
	Center Point   // 圆心坐标
}

// 点结构体
type Point struct {
	X float64 // x坐标
	Y float64 // y坐标
}

// 曲线方程
type CurveFunction func(x float64) float64

// 直线方程
type Line struct {
	Slope     float64 // 斜率
	Intercept float64 // 截距
}

/**
 * 1. 计算直线与非圆曲线的交点
 */
// 计算曲线上某点的函数值
func computeCurvePoint(f CurveFunction, x float64) float64 {
	return f(x)
}

// 计算直线上某点的函数值
func computeLinePoint(line Line, x float64) float64 {
	return line.Slope*x + line.Intercept
}

// 计算曲线和直线的交点横坐标
func findIntersectionX(curve CurveFunction, line Line, xMin, xMax, epsilon float64) float64 {
	xMid := (xMin + xMax) / 2
	yCurve := computeCurvePoint(curve, xMid)
	yLine := computeLinePoint(line, xMid)
	diff := yCurve - yLine

	if math.Abs(diff) <= epsilon || xMax-xMin <= epsilon {
		return xMid
	} else if diff > 0 {
		return findIntersectionX(curve, line, xMin, xMid, epsilon)
	} else {
		return findIntersectionX(curve, line, xMid, xMax, epsilon)
	}
}

// 计算直线和曲线的交点
func findIntersection(curve CurveFunction, line Line, xMin, xMax, epsilon float64) (float64, float64) {
	intersectionX := findIntersectionX(curve, line, xMin, xMax, epsilon)
	intersectionY := computeCurvePoint(curve, intersectionX)
	return intersectionX, intersectionY
}

/**
 * 2. 计算圆心和非圆曲线的公切线
 */
// 计算曲线上某点的斜率
func computeSlope(f CurveFunction, x float64) float64 {
	h := 1e-9 // 极小值
	return (f(x+h) - f(x-h)) / (2 * h)
}

// 计算切线的截距
func computeIntercept(f CurveFunction, x, y, slope float64) float64 {
	return y - slope*x
}

// 寻找切线的斜率和截距
func findTangentLineParameters(f CurveFunction, x, y float64) (float64, float64) {
	slope := computeSlope(f, x)
	intercept := computeIntercept(f, x, y, slope)
	return slope, intercept
}

// 寻找圆和曲线的公切线
func findCommonTangent(circle Circle, curve CurveFunction) ([]float64, []string) {
	tangentSlopes := make([]float64, 0)
	tangentEquations := make([]string, 0)

	for x := circle.Center.X - circle.Radius; x <= circle.Center.X+circle.Radius; x += 0.001 {
		// 计算圆上对应点的y坐标
		circleY := math.Sqrt(math.Pow(circle.Radius, 2) - math.Pow(x-circle.Center.X, 2))

		// 寻找切线的斜率和截距
		slope, intercept := findTangentLineParameters(curve, x, curve(x))

		// 判断切线是否与圆相切
		if math.Abs(slope*x-circleY+intercept) < 1e-6 {
			tangentSlopes = append(tangentSlopes, slope)

			// 构造切线方程的字符串表示
			equation := fmt.Sprintf("y = %.4f*x + %.4f", slope, intercept)
			tangentEquations = append(tangentEquations, equation)
		}
	}

	return tangentSlopes, tangentEquations
}

/**
 * 3. 求逼近直线
 */

// 计算两点之间的距离
func distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// 执行Douglas-Peucker算法逼近曲线
func douglasPeucker(points []Point, epsilon float64) []Point {
	if len(points) < 3 {
		return points
	}

	// 找到距离最大的点
	dmax := 0.0
	index := 0

	for i := 1; i < len(points)-1; i++ {
		d := distance(points[0], points[len(points)-1])
		if d > dmax {
			dmax = d
			index = i
		}
	}

	// 如果最大距离大于误差要求，则递归调用Douglas-Peucker算法
	if dmax > epsilon {
		results1 := douglasPeucker(points[:index+1], epsilon)
		results2 := douglasPeucker(points[index:], epsilon)

		// 将结果合并
		results1 = results1[:len(results1)-1]
		results1 = append(results1, results2...)

		return results1
	}

	return []Point{points[0], points[len(points)-1]}
}

func main() {
	// 定义圆
	circle := Circle{
		Radius: 1.0,
		Center: Point{
			X: 0.0,
			Y: 0.0,
		},
	}

	// 定义曲线方程
	curve := func(x float64) float64 {
		return math.Pow(x, 2) - 2*x + 1
	}

	// 设置允许加工误差
	epsilon := 0.1

	// 将曲线方程转换为离散的曲线点集
	var curvePoints []Point
	step := 0.1
	for x := 0.0; x <= 5.0; x += step {
		y := computeCurvePoint(curve, x)
		curvePoints = append(curvePoints, Point{X: x, Y: y})
	}

	// 执行Douglas-Peucker算法逼近曲线
	approximation := douglasPeucker(curvePoints, epsilon)

	// 定义直线方程
	line := Line{
		Slope:     2.0,
		Intercept: -1.0,
	}

	// 计算交点
	intersectionX, intersectionY := findIntersection(curve, line, -10.0, 10.0, 1e-6)

	// 输出结果
	fmt.Printf("交点坐标：(%.4f, %.4f)\n", intersectionX, intersectionY)

	// 寻找公切线的斜率和方程
	tangentSlopes, tangentEquations := findCommonTangent(circle, curve)

	// 输出结果
	fmt.Println("公切线的斜率：")
	for _, slope := range tangentSlopes {
		fmt.Println(slope)
	}

	fmt.Println("公切线的方程：")
	for _, equation := range tangentEquations {
		fmt.Println(equation)
	}

	// 输出逼近的直线段数据
	fmt.Println("逼近直线段数据：")
	for i := 0; i < len(approximation)-1; i++ {
		fmt.Printf("段 %d: (%.2f, %.2f) -> (%.2f, %.2f)\n", i+1, approximation[i].X, approximation[i].Y, approximation[i+1].X, approximation[i+1].Y)
	}

	// 输出逼近直线段的段数
	fmt.Printf("逼近直线段的段数：%d\n", len(approximation)-1)

	// 将逼近的直线段数据转化为CNC加工代码
	fmt.Println("\nCNC加工代码：")
	for i := 0; i < len(approximation)-1; i++ {
		fmt.Printf("G01 X%.2f Y%.2f\n", approximation[i+1].X, approximation[i+1].Y)
	}
}
