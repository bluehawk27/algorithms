

// 1. Let min = 0 and max = n-1.
// 2. If max < min, then stop: target is not present in array. Return -1.
// 3. Compute guess as the average of max and min, rounded down (so that it is an integer).
// 4. If array[guess] equals target, then stop. You found it! Return guess.
// 5. If the guess was too low, that is, array[guess] < target, then set min = guess + 1.
// 6. Otherwise, the guess was too high. Set max = guess - 1.
// 7. Go back to step 2.



function doSearch(array, targetValue) {
  var min = 0;
  var max = array.length - 1;
  var guess;
  while(min <= max ){
    guess = Math.floor((max + min)/2);

    if(array[guess] === targetValue) {
      return guess;}
    else if(array[guess] < targetValue){
      min = guess + 1;
    }
    else{
      max = guess - 1;
    }
  }
  return -1;
};

var primes = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
    41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97];

console.log(doSearch(primes, 73))
