import re

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

print(total_sum)