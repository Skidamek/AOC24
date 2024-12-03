import re
import time

start_time = time.time()

with open("input.txt") as file:
    text = "".join(line.strip() for line in file)

def get_matches(regex, text):
    matches = {}
    for match in re.finditer(regex, text):
        matches[match.start()] = match.group()
    return matches

matched_mul = get_matches(r"mul\(\d+,\d+\)", text)
matched_do = get_matches(r"do\(\)", text)
matched_dont = get_matches(r"don't\(\)", text)

disabled_pos = {}

for dont_start in sorted(matched_dont.keys()):
    end_range = max(matched_mul.keys())
    for do_start in sorted(matched_do.keys()):
        if do_start > dont_start:
            end_range = do_start
            break
    disabled_pos[dont_start] = end_range

disabled_coords = set()
for start, end in disabled_pos.items():
    disabled_coords.update(range(start, end + 1))

def is_disabled(position):
    return position in disabled_coords

total_sum = 0
for start, expression in matched_mul.items():
    if is_disabled(start):
        continue

    nums = map(int, re.findall(r"\d+", expression))
    result = 1
    for num in nums:
        result *= num

    total_sum += result

elapsed_time = (time.time() - start_time) * 1_000_000
print("Total sum:", total_sum)
print(f"Took: {elapsed_time:.3f}µs")