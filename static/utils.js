function isElementInView(elementID) {
    const element = document.getElementById(elementID);
    const rect = element.getBoundingClientRect();
    const windowHeight =
        window.innerHeight || document.documentElement.clientHeight;
    const windowWidth =
        window.innerWidth || document.documentElement.clientWidth;

    // check if element is completely in view
    const isInHorizontalView = rect.left >= 0 && rect.right <= windowWidth;
    const isInVerticalView = rect.top >= 0 && rect.bottom <= windowHeight;

    return isInHorizontalView && isInVerticalView;
}
