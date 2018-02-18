# Box Muller

from collections import Counter
from matplotlib import pyplot

import numpy as np


def box_muller():
    a, b = np.random.rand(2)
    ta = np.math.sin(2 * np.pi * b)
    tb = np.math.cos(2 * np.pi * b)

    t = np.math.sqrt(-np.log(a**2))
    return t * ta, t * tb


l = []
for i in range(100000):
    l += box_muller()

l2 = [sum(np.random.rand(40)) / 40 for _ in range(100000)]

fig = pyplot.figure()

ax = fig.add_subplot(2, 1, 1)
ax.hist(l, bins=1000)
ax.set_title('Box Muller')

bx = fig.add_subplot(2, 1, 2)
bx.hist(l2, bins=1000)
bx.set_title('Central limit theorem')

pyplot.show()