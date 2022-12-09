import os


def main():
    with open("day01.txt") as f:
        lines = [l.replace("\n", "") for l in f.readlines()]
    current_index = 0
    calories = {0: 0}
    for line in lines:
        if len(line) == 0:
            current_index += 1
            calories[current_index] = 0
        else:
            try:
                calorie = int(line.replace("\n", ""))
                calories[current_index] += calorie
            except Exception:
                print(f"<{line}> could not be parsed to an int")
    sorted_top = sorted(calories.values(), reverse=True)[:3]
    top_three = sum(sorted_top)

    print("Top three", sorted_top)
    print("Top three sum", top_three)
    print("Max value", sorted_top[0])


if __name__ == '__main__':
    main()
