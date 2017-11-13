def largest_prime_factorial(x)
  prime = x
  (2..Math.sqrt(x).to_i).each do |i|
     break if prime <= i
    prime /= i while (prime > i && prime % i == 0)
  end
  prime
end

p largest_prime_factorial(10000000000000000000)