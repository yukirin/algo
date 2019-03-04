# https://qiita.com/ganbarubaru/items/27e603acc16fee474fdd
import numpy as np
import matplotlib.pyplot as plt
import pandas as pd


def f(x):
    return 5 * x ** 4


def monte(N):
    a = 0
    b = 1

    x = (b - a) * np.random.rand(N)
    ans = np.average(f(x))
    return ans


startx = 1
endx = 1000
interval = 1

sample = np.arange(startx, endx + startx, interval)
m_int = pd.DataFrame((np.random.randn(sample.size, 2) + 2) * 0)
count = 0

for i in sample:
    m_int[0][count] = i
    m_int[1][count] = monte(i)
    count += 1

m_int = m_int.rename(columns={0: 'N', 1: 'ans'})
m_int.plot(x='N', figsize=(13, 8), title="y=5x^4")

plt.grid(which='major', color='black', linestyle='-')
plt.show()
