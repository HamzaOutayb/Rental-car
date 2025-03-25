import { useState } from 'react';
import styles from './addCarModal.module.css'; // Import the CSS module
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheck, faPlus, faTimes } from '@fortawesome/free-solid-svg-icons';

export default function AddCarModal({ toggleModal }) {
  const [carName, setCarName] = useState('');
  const [description, setDescription] = useState('');
  const [rentPrice, setRentPrice] = useState('');
  const [images, setImages] = useState([]);
  const [brand, setBrand] = useState('');
  const [type, setType] = useState('');
  const [contact, setContact] = useState('');
  const [conditions, setConditions] = useState([]);
  const [condition, setCondition] = useState('');

  // Handle Image Upload
  const handleImageUpload = (e) => {
    const files = Array.from(e.target.files);
    if (files.length > 0) {
      setImages((prevImages) => [...prevImages, ...files]);
    }
  };

  // Remove Image
  const removeImage = (index) => {
    setImages(images.filter((_, i) => i !== index));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission
  };

  const addCondition = () => {
    setConditions([...conditions, condition]);
    setCondition('');
  }

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modal}>
        <span onClick={toggleModal} className={styles.close}>
          <FontAwesomeIcon icon={faTimes} />
        </span>
        {/* <h2>Add New Car</h2> */}
        <form onSubmit={handleSubmit} className={styles.form}>

          {/* First Row: Car Name, Rent Price, Description */}
          <div className={styles.formRow}>
            <input
              className={styles.input}
              type="text"
              value={carName}
              onChange={(e) => setCarName(e.target.value)}
              placeholder="Car Name"
              required
            />
            <input
              className={styles.input}
              type="text"
              value={rentPrice}
              onChange={(e) => setRentPrice(e.target.value)}
              placeholder="Rent Price"
              required
            />
          </div>
          <textarea
            className={styles.textarea}
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Description"
            required
          ></textarea>


          <div className={styles.formRow}>
            <select className={styles.select} value={brand} onChange={(e) => setBrand(e.target.value)} required>
              <option value="">Select Brand</option>
              <option value="brand1">Brand 1</option>
              <option value="brand2">Brand 2</option>
            </select>
            <select className={styles.select} value={type} onChange={(e) => setType(e.target.value)} required>
              <option value="">Select Type</option>
              <option value="sedan">Sedan</option>
              <option value="suv">SUV</option>
            </select>
            <select className={styles.select} value={contact} onChange={(e) => setContact(e.target.value)} required>
              <option value="">Select Contact</option>
              <option value="hamza.id">hamza</option>
              <option value="mohamed.id">mohamed</option>
            </select>
          </div>
          {/* conditions preview */}
          <div className={styles.conditionsCon}>
            {
              conditions.map((el, i) => {
                return (
                  <div key={i}>
                    <FontAwesomeIcon icon={faCheck} /> {'  '}
                    {el}
                  </div>
                )
              })
            }
          </div>
          <span className={styles.condInputCont}>
            <input
              className={`${styles.conditionInput}`}
              type="input"
              value={condition}
              onChange={(e) => setCondition(e.target.value)}
              placeholder="Conditions"
              required
            />

            <FontAwesomeIcon onClick={addCondition} className={styles.add} icon={faPlus} />
          </span>

          {/* Image Upload */}
          <input className={styles.fileInput} type="file" accept="image/*" multiple onChange={handleImageUpload} />

          {/* Image Preview Section */}
          <div className={styles.imagePreviewContainer}>
            {images.map((image, index) => (
              <div key={index} className={styles.imageWrapper}>
                <img src={URL.createObjectURL(image)} alt={`Uploaded ${index}`} className={styles.imagePreview} />
                <button type="button" className={styles.removeImageButton} onClick={() => removeImage(index)}>
                  <FontAwesomeIcon icon={faTimes} />
                </button>
                {index === 0 && <span className={styles.primaryLabel}>Primary</span>}
              </div>
            ))}
          </div>

          <button className={styles.button} type="submit">Add Car</button>
        </form>
      </div>
    </div>
  );
}
