,          Read first character
[          Start loop until input is zero (EOF)
  > ,      Move to next cell and read next character
]          End loop
<          Move back to last valid character
[          Loop backwards to print characters
  .        Print current character
  <        Move to previous character
]          End loop
