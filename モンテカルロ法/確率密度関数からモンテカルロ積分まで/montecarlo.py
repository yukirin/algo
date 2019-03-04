# https://qiita.com/Ushio/items/0040b3c74a480c46c80c
# f(x) = exp(-x^2/0.5)
# p(x) = 1 - 1/2x
# f(x) {0 <= x <= 2}の積分を近似
# 積分の真値 0.626617

from numpy import random
import math


def f(x):
    return math.exp(-x * x / 0.5)


def p(x):
    return 1.0 - 0.5 * x


def inverse_of_cdf(x):
    return 2.0 - 2.0 * math.sqrt(1.0 - x)


total = 0
N = 60000

for i in range(N):
    if i != 0 and i % 1000 == 0:
        ans = total / i
        diff = abs(ans - 0.626617)
        print(f"diff = {diff}")

    r = random.rand()
    x = inverse_of_cdf(r)
    value = f(x) / p(x)
    total += value
