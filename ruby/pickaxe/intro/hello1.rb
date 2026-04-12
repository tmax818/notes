def say_hello_goodbye(name)
    result = "I don't know why you say goodbye, " + name + ", I say hello"
    return result
end

# call the method​
puts say_hello_goodbye "john"
puts say_hello_goodbye("Ringo")