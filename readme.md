### 圆函数曲线等误差直线段逼进算法及CNC代码的生成

#### 课程设计的内容

非圆曲线是指二次或者二次以上的除圆曲线以外的初等函数所表示的曲线方程，如抛物线，双曲线，螺旋线，以及其他等曲线。非圆曲线是数控编程中经常用到的一类复杂轮廓曲线，对非圆曲线进行数控编程的关键点在于如何提高数控程序的质量，以避免数控加工程序段数量过多。对非圆曲线的数学处理，其运算量和计算复杂性，是手工编程不能胜任的。目前流行的商业化CAD软件并未考虑数控加工代码的优化问题。本课程设计的目的合意义就在干针对数控加工需要而商务CAD系统却未提供的一些功能，要求学生通过创新开发一套自主版权软件工具来满足实际的需求。

#### 课程设计的要求与数据

本课程设计的要求是:对用户任意给定的一条非圆轮廓曲线及允许加工误差，用最少的直线段来逼进给定的非圆曲线轮廓，以使得数控加工程序段数最少。具体的要求分为如下几个部分:

1. 设计一个友好的用户界面，使得用户能够完成曲线的输入，轮廓曲线显示，显示逼进曲线段数，显示数控代码的生成。最好有(但不作考核要求)查看图形的各种操作，如显示全图，放大、缩小，局部窗口的放大、缩小，上下左右的移动等。

2. 设计并编写一个软件算法，它能求出圆和非圆曲线的公切线。

3. 设计并编写一个软件算法，它能求出直线和非圆曲线的交点。

4. 将求解的逼进直线段及其段数显示在屏幕上。

5. 将逼进的直线段数据转化为CNC加工代码。