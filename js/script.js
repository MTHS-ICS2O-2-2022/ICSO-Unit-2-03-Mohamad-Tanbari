// Copyright (c) 2023 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: Sep 2023
// This file contains the JS functions for index.html

/*
* This function collects the user's input and displays a combined result.
*/
function buttonClicked() {
  // input
  const streetName = document.getElementById('street-name').value
  const streetNumber = parseInt(document.getElementById('street-number').value)

  // output
  document.getElementById('address').innerHTML = 
    "Your address is: " + streetName + " " + streetNumber + "."
}
