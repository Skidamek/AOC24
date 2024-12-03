import re
import time

start_time = time.time()

with open("input.txt") as file:
    text = "".join(line.strip() for line in file)

matchedParts = re.findall(r"mul\(\d+,\d+\)", text)

total_sum = 0

for match in matchedParts:
    nums = re.findall(r"\d+", match)
    result = 1
    for n in nums:
        result *= int(n)

    total_sum += result

elapsed_time = (time.time() - start_time) * 1_000_000
print("Total sum:", total_sum)
print(f"Took: {elapsed_time:.3f}µs")