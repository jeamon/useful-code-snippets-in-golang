package main

import ("fmt"; "strings")


/*
// 
//   During a coding interview - I was tasked to craft below string re-ordering
//   
//   ex: For input "LLKMLKRTYSHKFFKLMFEWPRLGFSPRT" output should be "LLLLLKKKKMMRRRTTYSSHFFFFEWPPG"
//   
//   FYI - Order by first occurance of the character
// 
*/

// take string as input and return string
func arrange(inputString string) string {

  mapOccurrences := make(map[string]int)
  inputListWithoutDuplicates := make([]string, 0)
  var result string
  var charIntoString string
  
  for _, char := range inputString {
    
    // char is rune type - convert into string type
    charIntoString = string(char)
    
    if occurrence, exists := mapOccurrences[charIntoString]; exists {
      // increment that character occurence
      mapOccurrences[charIntoString] = occurrence + 1
    
    } else { // initialize occurence and add to list
      
      mapOccurrences[charIntoString] = 1
      inputListWithoutDuplicates = append(inputListWithoutDuplicates, charIntoString)
    
    }
  }
  
  // loop over new list without duplicate to build result string
  // for each character - repeat it with its occurrences number
  for _, character := range inputListWithoutDuplicates {
    result = result + strings.Repeat(character, mapOccurrences[character])
  }
  
  return result
}

// entry point
func main() {

  str := "LLKMLKRTYSHKFFKLMFEWPRLGFSPRT"
  fmt.Println(arrange(str))
  
}