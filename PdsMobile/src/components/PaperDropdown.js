import React, {useState} from "react";
import DropDown from "react-native-paper-dropdown";

const PaperDropdown = ({
  listValue = [],
  label,
  value,
  setValue,
  mode = "outlined",
  ...props
}) => {
  const [showDropDown, setShowDropDown] = useState(false);
  return (
    <DropDown
      {...props}
      label={label}
      mode={mode}
      visible={showDropDown}
      showDropDown={() => setShowDropDown(true)}
      onDismiss={() => setShowDropDown(false)}
      value={value}
      setValue={setValue}
      list={listValue}
    />
  );
};

export default PaperDropdown;
