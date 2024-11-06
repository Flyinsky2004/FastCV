import dayjs from 'dayjs';

function dateToString(date) {
    const dayjsDate = dayjs.isDayjs(date) ? date : dayjs(date);
    if (!dayjs.isDayjs(dayjsDate)) {
        throw new Error('Input must be a dayjs object');
    }
    return dayjsDate.format('YYYY-MM-DD');
}
export default dateToString;