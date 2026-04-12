
doubleMe x = x + x

doubleUs x y = x * 2 + y * 2
 
doubleSmallNumber x = if x > 100 -- if is an expression not a statement
    then x      
    else x*2    -- else is mandatory