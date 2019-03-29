var g1 = {
    k: 1,
    n: "中横通讯",
    price1: 35.506,
    price2: 38.600,
    store1: 600,
    store2: 600,
    charge1: 1856.47,
    charge2: 0.08714
};
var g2 = {k: 2, n: "震惊股份", price1: 10.452, price2: 9.58, store1: 2400, store2: 2400, charge1: -2093, charge2: -0.08343};
var g3 = {k: 3, n: "长油", price1: 2.392, price2: 3.270, store1: 4300, store2: 4300, charge1: 3774.21, charge2: 0.36706};
var go1 = {zongzichan: 61154.23, fudong: 3537.68, dangriyingkui: 5507, zongshizhi: 60213, keyong: 940.63, kequ: 940.63};
var go2 = {zongzichan: 61154.23, fudong: 3537.68, dangriyingkui: 1, zongshizhi: 1, keyong: 1, kequ: 1};
var selllog = [{k: 2, count: 2400, price: 9.33}, {k: 3, count: 4300, price: 3.34}];
var buylog = [];
var chicang = [g1, g2, g3];
var sel = {n: 2, count: 1, price: 1}

function cal() {
    for (var i = 0; i < selllog; i++) {
        const fg = selllog[i];
        const cg = getChicang(f.k)
        cg.store1-=fg.count
        cg.price1

    }
}

function getChicang(k) {
    return chicang.find(value => value.k == k)
}