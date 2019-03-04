import numpy as np
import matplotlib.pyplot as plt


a = 0
b = 2
M = 2


def f(x):
    return np.sin(x)


def rejection_sampling(N):
    x = (b - a) * np.random.rand(N)
    u = np.random.rand(N)
    accept_x = (-1) * np.ones(N)

    for i in range(N):
        if M * u[i] < f(x)[i]:
            accept_x[i] = x[i]

    return accept_x


N = 10000
result = rejection_sampling(N)

# f(x)
x = np.arange(a, b, 0.01)
plt.plot(x, (N / 200) * f(x))

# histogram
plt.hist(result, bins=100, range=(a, b))
plt.grid(which='major', color='black', linestyle='-')
plt.show()
