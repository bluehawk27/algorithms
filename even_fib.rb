def even_fib(max_num)
  even_sum = 0
  val1 = 1
  val2 = 1
  temp_sum = val1 + val2
  while temp_sum <= max_num
    if temp_sum % 2 == 0
      even_sum += temp_sum
    end
     val1 = val2
     val2 = temp_sum
     temp_sum = val1 + val2
  end
  even_sum
end

p even_fib(4000000)