<?php

$handle = fopen("php://stdin", "r");
fscanf($handle, "%d %d", $length, $rotations);
$a_temp = fgets($handle);
$numbers = explode(" ", $a_temp);

$new_array = array();
foreach ($numbers as $key => $value) {
  $new_key = $key - $rotations;
  if ($new_key >= 0) {
    $new_array[$new_key] = $value;
  } else {
    $new_key = $new_key + $length;
    $new_array[$new_key] = $value;
  }
}

ksort($new_array);
$numbers_to_string = implode($new_array,' ');
print $numbers_to_string;
?>
