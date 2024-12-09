array = list(range(101))  # 0 to 100 inclusive

# Brute force using for loop to solve FizzBuzz;
for x in array:
    if x % 3 == 0 and x % 5 == 0:
        print("FizzBuzz!")
    elif x % 3 == 0:
        print("Fizz!")
    elif x % 5 == 0:
        print("Buzz!")
    else:
        print(x)

# Recursive example of solving FizzBuzz;


def recurse(arr, n):
    if n >= len(arr):
        print("Finished recursion!")
        return

    if n % 3 == 0 and n % 5 == 0:
        print("FizzBuzz!")
    elif n % 5 == 0:
        print("Buzz!")
    elif n % 3 == 0:
        print("Fizz!")
    else:
        print(arr[n])

    recurse(arr, n + 1)


# Start recursion
recurse(array, 0)
