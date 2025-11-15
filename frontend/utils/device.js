export function isAndroidWebView() {
    const ua = navigator.userAgent || '';
    console.log('=== DEVICE DETECTION ===');
    console.log('UserAgent:', ua);

    const isAndroid = /Android/i.test(ua);
    const hasWebView = /wv|WebView/i.test(ua);
    const isChromeBrowser = /Chrome\//.test(ua) && !/wv|WebView/.test(ua);

    // Android WebView = Android + WebView маркер + НЕ Chrome браузер
    const result = isAndroid && hasWebView && !isChromeBrowser;

    console.log('Is Android:', isAndroid);
    console.log('Has WebView:', hasWebView);
    console.log('Is Chrome Browser:', isChromeBrowser);
    console.log('Final - Is Android WebView:', result);

    return result;
}