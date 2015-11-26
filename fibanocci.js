//nth fibanocci
function Fibanocci(n) {
  var result = [0,1]

  for(var i = 1; i < n; i++){
    result[i+1] = result[i] + result[i-1]
  }
  return result[n];
}

console.log(Fibanocci(10))
//55