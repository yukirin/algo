from numpy import random
import math

L = 2
d = L * 2

count = 100000
overlapped = 0

for _ in range(count):
    thita = random.rand() * math.pi
    y = random.rand() * (d / 2)

    if y <= (L / 2) * math.sin(thita):
        overlapped += 1

p = overlapped / count

print(f"p = {p}, pi = {1 / p}")
