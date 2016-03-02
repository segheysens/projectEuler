/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

var factors = [];
var number = 10086647;

for(var i = 3; i <= number; i += 2 ){
  if (number % i == 0) {
    factors.push(i);
    factors.push(number/i);
    number = number / i;
  }
}

number = 8462696833;
var limit = factors.length;

// for(var index = 0; index < limit; index++ ){
//   console.log("The current factor is " + factors[index]);
//   for(var i = 2; i <= number; i += 1 ){
//     if (factors[index] % i == 0 && factors[index] != i) {
//       console.log("The current index factors[" +index+ "] contains the factor " + factors[index] + " and we're dividing: " +factors[index]+ "/"+i+".");
//       console.log("This potentially new prime factor is " + i);
//       console.log("Here's the array before we may insert anything: [" + factors + "]");
//       console.log("factors.indexOf(i) = "+ factors.indexOf(i));
//       if(factors.indexOf(i) == -1){factors.push(number/i);console.log("We decided to add "+(number/i))}
//       number = number / i;
//       console.log("Here's the array after we insert anything: [" + factors + "]\n");
//       break;
//     }
//   }
// }

console.log(factors)
