fn main() {
    let array: Vec<i32> = (0..101).collect();
    let slice = &array[..];
    // Loop solution for FizzBuzz!;
    for n in slice.iter().cloned() {
        if n % 3 == 0 && n % 5 == 0 {
            println!("FizzBuzz!");
        }
        if n % 5 == 0 {
            println!("Buzz!")
        }
        if n % 3 == 0 {
            println!("Fizz!")
        }
        println!("{:#?}", n);
    }
    //Recursive solution for fizzbuzz;
    fn recurse(arr: &[i32], n: u32) {
        if n as usize == arr.len().try_into().unwrap() {
            println!("Finished Recursing!");
            return;
        }
        if arr[n as usize] % 15 == 0 {
            println!("FizzBuzz!")
        }
        if arr[n as usize] % 5 == 0 {
            println!("Buzz!");
        }
        if arr[n as usize] % 3 == 0 {
            println!("Fizz!")
        }
        println!("{:#?}", arr[n as usize]);
        return recurse(&arr, n + 1);
    }
    recurse(&array, 0)
}
