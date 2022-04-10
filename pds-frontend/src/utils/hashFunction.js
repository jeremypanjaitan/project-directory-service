import md5 from "md5";
const hashFunction = (text, _md5 = md5) => {
  return _md5(text);
};
export default hashFunction;
