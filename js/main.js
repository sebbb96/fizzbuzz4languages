const array = Array.from(Array(101).keys());

//Brute force FizzBuzz example with a for loop;

for (let i = 0; i < array.length; i++) {
  if (array[i] % 3 === 0 && array[i] % 5 === 0) {
    console.log("Fizz Buzz!");
  }
  if (array[i] % 3 === 0) {
    console.log("Fizz!");
  }
  if (array[i] % 5 === 0) {
    console.log("Buzz!");
  }
  console.log(array[i]);
}

//Recursive example for FizzBuzz Solving;

function recurse(arr, n) {
  if (n === arr.length) {
    console.log("Finished recursion", arr[n]);
    return;
  }
  if (arr[n] % 3 === 0 && arr[n] % 5) {
    console.log("FizzBuzz!");
  }
  if (arr[n] % 5 === 0) {
    console.log("Buzz!");
  }
  if (arr[n] % 3 === 0) {
    console.log("Fizz!");
  }
  console.log(arr[n]);
  return recurse(arr, n + 1);
}

recurse(array, 0);
