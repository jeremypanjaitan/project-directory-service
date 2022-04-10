import moment from "moment";
const formatDate = (date, _moment = moment) => {
  return _moment(date).format("DD MMMM YYYY");
};

export default formatDate;
