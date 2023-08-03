if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}
const go = new Go();  // wasm_exec.js で定義されたGoオブジェクト
let wasmModule;

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then(result => {
        wasmModule = result.instance;
        go.run(wasmModule);  // wasm_exec.js のランタイムを起動
    });

function loadImage() {
    const input = document.getElementById('image-upload');
    const file = input.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = function (e) {
        const image = new Image();
        image.onload = function () {
            const canvas = document.getElementById('image-canvas');
            canvas.width = image.width;
            canvas.height = image.height;
            const ctx = canvas.getContext('2d');
            ctx.drawImage(image, 0, 0);
        };
        image.src = e.target.result;
    };
    reader.readAsDataURL(file);
}
    
function getImageData() {
    const canvas = document.getElementById('image-canvas');
    const ctx = canvas.getContext('2d');
    const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);
    return new Uint8Array(imageData.data.buffer);
}

function applyGrayscale() {
    const strength = document.getElementById('grayscale-strength').value;
    const imageData = getImageData();
    const resultData = new Uint8Array(applyGrayscale(imageData, strength));
    renderImage(resultData);
}

function resizeImage() {
    const newWidth = document.getElementById('width').value;
    const newHeight = document.getElementById('height').value;
    const imageData = getImageData();
    const resultData = new Uint8Array(resize(imageData, newWidth, newHeight));
    renderImage(resultData);
}

function rotateImage() {
    const angle = document.getElementById('rotation-angle').value;
    const imageData = getImageData();
    const resultData = new Uint8Array(rotate(imageData, angle));
    renderImage(resultData);
}

function flipHorizontal() {
    const imageData = getImageData();
    const resultData = new Uint8Array(flipHorizontal(imageData));
    renderImage(resultData);
}

function renderImage(byteArray) {
    const canvas = document.getElementById('image-canvas');
    const ctx = canvas.getContext('2d');
    const imageData = new ImageData(new Uint8ClampedArray(byteArray), canvas.width, canvas.height);
    ctx.putImageData(imageData, 0, 0);
}
