//assuming all inputs are positiv integers.  Non integers will create infinite recursion

var factorial = function(n) {

    if (n === 0) {
        return 1;
    }else{
        return n * factorial(n-1);
    }
};

console.log(factorial(10))
//3628800