import { formatInTimeZone } from "date-fns-tz";

const dateFormat = (
  date: Date | string,
  format: string = "dd MMM, yyyy",
): string => {
  if (!date) {
    return "";
  }
  return formatInTimeZone(date, "America/New_York", format);
};

export default dateFormat;

