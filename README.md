## Latitude Code Test

This is my solution for [this code test](https://gist.github.com/jonog/54e46b5b1200758d222e3c4cf61baaa6). Essentially the problem as I see it is we are looking for is the max difference within an array of integers where the conditions are such that the larger of the two integers used to calculate the diff must be positioned after the smaller integer within our array. So the maximum difference in `[10,8,6,4,2]` would be 0, but in `[2,4,6,8,10]` it would be 8 (10 - 2).

The following assumptions and their justifications were made;

* We don't need to restrict the array size to exactly 360 (6 hours of trading times 60 minutes per hour), as the solution is the same so long as the array length is >= 2
* We do not need to know the exact times that we needed to buy and sell, only that we must buy before we sell. 
* The spec asks for an *efficient function*, rather than a fully featured program, so the scope of the program will be limited to executing the function against a hard coded array.
* The term *efficient function* doesn't necessarily imply the **most** efficient function (though it is something to always strive for, within reason). As such I'm going to assume that anything that anything O(n) or better is acceptable

### Solution

So obviously a straight forward way to solve this is with an outer loop that traverses the elements of the array, that contains an inner loop that compares the inner loop value with the outer loop value to find our maximum profit. This however will run in O(n²) time, which is not efficient.

Instead we can use a pretty well known algorithm that runs in O(n). 
* Initialize our max diff with a[1] - a[0]
* Initialize our min with a[0]
* Traverse the array starting from the second element (we already have a diff from the first), and check at each index if the value - min > max diff, and if the value < min
* If these conditions are met, update the max diff and min values accordingly

Once the array has been traversed, our max diff value should be the largest difference within the array where the index of the larger value occurs subsequent to the smaller value, and our min should be the smallest value within the array. Since we don't need the min value other than to calculate the max diff, we don't need to store the min value relative to the max diff, so this is acceptable.

### Test Scenarios

Since I don't work in finance, I don't know what a full list of scenarios looks like, so we will limit our tests to the following stock price patterns, which should be pretty all encompassing anyway.

#### Linear Rise
```
     /
    /
   /
  /
 /
/
```

#### Linear Fall
```
\
 \
  \ 
   \
    \
     \
```

#### Falling With Spike
```
\_
  \_    /\
    \  /  \
     \/    \
            \
             \_
               \
```

#### Rising With Spike
```
      /\
     /  \    /
    /    \  /
  _/      \/
 /         
/
```

#### Flat With Spike
```
               /\
              /  \  
             /    \
            /      \
_____/\____/        \____
```

#### Multiple Spikes
```

            /\        /
    /\    _/  \      /
 /\/  \  /     \__  /
/      \/         \/
```