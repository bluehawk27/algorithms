//input: takes in a string "abcdefghijklmnop"
//output: returns it in a new string from reverse variable "ponmlkjihgfedcba"



function sRev(input){
  var reverse = ""

  function reversal(index){
    if(input.length === reverse.length){
      return;
    }
    reverse = reverse + input[index]
    reversal(index - 1)
  }

  reversal(input.length -1);
  return reverse;
}

console.log(sRev("abcdefghijklmnop"))



//reverse string in place
function rString(input){
  strArr = input.split('')

  function reverse(start, end) {
    if (start >= end){
      return;
    }
    temp = strArr[start]
    strArr[start] = strArr[end]
    strArr[end] = temp
    reverse(start + 1, end - 1)
  };

  reverse(0, input.length - 1)
  return strArr.join('')
};

console.log(rString("Alberto"))
//output: otreblA