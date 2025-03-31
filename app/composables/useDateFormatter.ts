import {format} from "date-fns";

export const useDateFormatter = () => {
    const formatDate = (date: string) => format(date, "do MMMM yyyy");

    return {formatDate};
};