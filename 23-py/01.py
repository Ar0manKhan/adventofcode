numberStrings = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

def printNumbersInLine():
    with open('input.txt', 'r') as f:
        for lx, line in enumerate(f):
            strippped_line = line.strip()
            for i in range(len(strippped_line)):
                # if stripped_line[i] is digit, print it
                if strippped_line[i].isdigit():
                    print(strippped_line[i], end='')
                    continue
                substringline = strippped_line[i:]
                for num, digit in enumerate(numberStrings):
                    if substringline.startswith(digit):
                        print(num, end='')
                        break
            print()



if __name__ == '__main__':
    printNumbersInLine()
