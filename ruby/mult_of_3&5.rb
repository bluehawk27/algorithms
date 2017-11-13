def multiples(number)

  unless number.is_a? Integer
    return "Input must be a number."
  end

  unless number > 0
    return "Input must be greater than 0."
  end

  sum = 0
  count = 1
  while count < number
    if count % 3 == 0 || count % 5 == 0
      sum += count
    end
    count+=1
  end
  sum
end

p multiples(1000)