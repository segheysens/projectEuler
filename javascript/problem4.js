/*
A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

// Function should return TRUE is the number is a palindrome
function isPalindrome(number){
  var tempString = (number).toString();
  // Loop through the number (index from the )
  for(var index = 0; index < (tempString.length/2); index += 1){
    // return false if the digit at index != its mirror digit
    if(tempString.charAt(index) !== tempString.charAt(tempString.length-index-1)){
      return false;
    }
  }
  // if the loop didn't return false, number is a palindrome
  return true;
}

var i, j, largestPal = 0;

/* To find the largest palindrome, it makes sense to start with the largest
possible product of three-digit numbers and work out way down! 999*999
*/
for(i = 999; i >= 0; i -= 1){
  for(j = 999; j >= 0; j -= 1){
    var prod = i * j;
    if(isPalindrome(prod) && prod > largestPal){
      largestPal = prod;
      break;
    }
  }
}

console.log(largestPal);
