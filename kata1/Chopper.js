function chop(value, values) {
    debugger;
    if (values.length === 0) {
        return -1;
    }

    var position = values.length,
        middle = values.length / 2,
        a = values,
        b = values;

    while (true) {
        if (value < a[a.length - 1]) {
            values = values.slice(0, middle);
        } else if (value < b[b.length - 1]) {
            values = values.slice(middle);
        } else if (value == a[0] || value == b[0]) {
            return position;
        }

        if (a == [] && b == []) return -1;

        middle /= 2;
        position -= middle;
        a = values.slice(0, middle);
        b = values.slice(middle);
    }
}
