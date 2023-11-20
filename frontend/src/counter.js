export function setupCounter(element) {
    let counter = 0
    function setCounter(count) {
        counter = count
        element.innerHTML = `count is ${counter}`
    }
    element.addEventListener('click', (e) => {
        setCounter(counter + 1)
        console.log(e.target)
    })
    setCounter(0)
}
