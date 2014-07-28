function NotInArray_ReturnsNegativeOne() {
    console.assert(-1 === chop(3, [1]));
}

function SingleValueArray_ValueFound_ReturnsZero() {
    console.assert(0 === chop(1, [1]));
}

function MultipleValueArray_ValuetFound_ReturnsZero() {
    console.assert(0 === chop(1,[1, 3, 5]));
}

function MultipleValueArray_ValueFound_ReturnsTwo() {
    console.assert(2 === chop(5, [1, 3, 5, 7]));
}

NotInArray_ReturnsNegativeOne();
SingleValueArray_ValueFound_ReturnsZero();
MultipleValueArray_ValuetFound_ReturnsZero();
MultipleValueArray_ValueFound_ReturnsTwo();
