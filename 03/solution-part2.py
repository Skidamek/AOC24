import re

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

print("mul:", matched_mul)
print("do:", matched_do)
print("dont:", matched_dont)

disabled_pos = {}

for dont_start in sorted(matched_dont.keys()):
    end_range = max(matched_mul.keys())
    for do_start in sorted(matched_do.keys()):
        if do_start > dont_start:
            end_range = do_start
            break
    disabled_pos[dont_start] = end_range

print("disabled_pos:", sorted(disabled_pos.items()))

disabled_coords = set()
for start, end in disabled_pos.items():
    disabled_coords.update(range(start, end + 1))

print("disabled_coords:", sorted(disabled_coords))

def is_disabled(position):
    return position in disabled_coords

total_sum = 0
for start, expression in matched_mul.items():
    print("mul", start)
    if is_disabled(start):
        print("disabled")
        continue

    print("enabled")
    nums = map(int, re.findall(r"\d+", expression))
    result = 1
    for num in nums:
        result *= num

    total_sum += result
    print("result:", result)

print(total_sum)
