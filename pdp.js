function changeImage(imageSrc) {
    document.getElementById('mainImage').src = imageSrc;
}

function openImageModal() {
    document.getElementById('popupImage').src = document.getElementById('mainImage').src;
    document.getElementById('imageModal').style.display = 'flex';
}

function openContactModal() {
    document.getElementById('contactModal').style.display = 'flex';
}

function closeModal(modalId) {
    document.getElementById(modalId).style.display = 'none';
}