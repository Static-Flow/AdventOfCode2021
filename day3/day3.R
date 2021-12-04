library(bit)
library(binaryLogic)
library(stringr)

#This is my solution for Day 3 of AdventOfCode2021 in R.
#https://adventofcode.com/2021/day/3
# This is likely not performant R code at ALL. It taught myself enough
# R to solve these but I'm almost certain this could be better. Likely
# a lot of wasted type conversions or snazzier list comprehessions I could use


#This converts a vector of bits to an integer
#Lifted from here https://stackoverflow.com/a/25411493/1827657
bitsToInt<-function(x) {
  packBits(rev(c(rep(FALSE, 32-length(x)%%32), as.logical(x))), "integer")
}

# Solution to Part 1 of Day 3
# A key performance improvement is to note that the Epsilon rate is just the Bitwise NOT of the Gamma rate
# so we only need to calculate the Gamma value. The way I decided to do that was keep a "bitmask" of sorts for
# each bit of the binary input and if it was a 1, add 1 to the corresponding bitmask value at the matching index
# or subtract 1 if it was 0. After all the values have been computed, any negative value in the bitmask means
# that "column" across all inputs had more 0's than 1's. Any positive value in the bitmask means more 1's than 0's.
# After replacing the negative values with 0, and positive values with 1 the bitmask is now the binary representation
# of the Gamma value. Once you have that, just bitwise NOT Gamma to get Epsilon and you're home free.
part1<-function() {
  # we setup a mask for all 12 bits in each binary input
  calcs <- c(0,0,0,0,0,0,0,0,0,0,0,0)
  #read in the file. Does R have streaming I/O? Dunno but this loads it all into memory
  to.read <- file("input","r")
  res <- readLines(to.read)
  # for each input line
  for (input in res) {
    # split the binary number (formatted as a string) into individual characters
    numberstring_split <- strsplit(input, "")[[1]]
    # for each input we reset out spot in the global calcs vector
    calcs_index <- 1
    # for each character in the split apart input
    for (number in numberstring_split) {
      # convert the bit character to an integer
      str_to_int <- base::strtoi(number, base = 10)
      # Bitwise AND it with 1 to see if it's set
      bit_set_check <- bitwAnd(1,str_to_int)
      if(bit_set_check == 0) {
        # if it isn't set subtract 1 from our global calcs vector
        calcs[calcs_index] <- calcs[calcs_index]-1
      } else {
        # if it is, add 1 to the global calcs vector
        calcs[calcs_index] <- calcs[calcs_index]+1
      }
      # advance global calc vector index
      calcs_index <- calcs_index+1
    }
  }
  # this comprehensions replaces any calcs value greater or equal to 0 with 1
  # and any negative value with 0
  calcs <- sapply(calcs, function(e) {
    if(e >= 0) {
      1
    } else {
      0
    }
  })
  # we convert the calcs global vector from strings into a single integer
  gamma <- base::strtoi(paste(calcs, collapse = ""),base=2)
  # we convert that integer into a binary form perform a bitwise NOT to flip all the bits
  # then return it to an integer
  epsilom <- bitsToInt(!as.bit(as.binary(gamma)))
  # multiply the values and return it
  return(gamma * epsilom)
}

# Solution to Part 2 of Day 3
# See recursiveCheck() for algorithm details
part2<-function () {
  # read in the input
  to.read <- file("input","r")
  res <- readLines(to.read)
  # send the initial values to the recursive solver
  rounds <- recursiveCheck(res,res,1)
  close(to.read)
  return(rounds[1]*rounds[2])
}

# recursive solver for Part 2 of Day 3. There is probably an interative solution but I couldn't think
# of how to do it AND learn the R way so I came up with this recursive version that works!
recursiveCheck<-function(ox_inputs_vector, co_inputs_vector, count) {
  ox_calcs <- replicate(length(ox_inputs_vector),0)
  ox_reduced <- NULL
  co_calcs  <- replicate(length(co_inputs_vector),0)
  co_reduced <- NULL
  TopSorted <- NULL
  # for each oxygen input
  if(length(ox_inputs_vector) > 1) {
    for(index in seq_along(ox_inputs_vector)) {
      ox_calcs[index] <-  base::strtoi(strsplit(ox_inputs_vector[index], "")[[1]][count],base = 10)
    }
    sorted_table <- sort(table(ox_calcs),decreasing=TRUE)

    if(sorted_table[1] == sorted_table[2]) {
      TopSorted <- 1
    } else {
      TopSorted <- base::strtoi(names(sorted_table)[1],base = 10)
    }
    ox_calcs <- sapply(ox_calcs, function(e) {
      if(e == TopSorted) {
        1
      } else {
        0
      }
    })
    ox_inputs_vector <- ox_inputs_vector[as.logical(ox_calcs)]
  }
  if(length(co_inputs_vector) > 1) {
    for(index in seq_along(co_inputs_vector)) {
      co_calcs[index] <-  base::strtoi(strsplit(co_inputs_vector[index], "")[[1]][count],base = 10)
    }
    sorted_table <- sort(table(co_calcs),decreasing=TRUE)
    if(sorted_table[1] == sorted_table[2]) {
      TopSorted <- 0
    } else {
      TopSorted <- base::strtoi(names(sorted_table)[2],base = 10)
    }
    co_calcs <- sapply(co_calcs, function(e) {
      if(e == TopSorted) {
        1
      } else {
        0
      }
    })
    co_inputs_vector <- co_inputs_vector[as.logical(co_calcs)]
  }

  if(length(ox_inputs_vector) == 1 && length(co_inputs_vector) == 1) {
    return(c(base::strtoi(paste(ox_inputs_vector[1], collapse = ""),base=2),base::strtoi(paste(co_inputs_vector[1], collapse = ""),base=2)))
  } else {
    recursiveCheck(ox_inputs_vector,co_inputs_vector,count+1)
  }

}

part1()
part2()



