import { faCaretDown, faPlus } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import AddCarModal from './addCarModal';
import styles from './cars.module.css'
import { useState } from 'react';
// import Image from 'next/image';

const types = ['SUV', 'VAN', 'LUX']
const brands = [
    { name: 'BMW', logoPath: '/bww-logo-1963.webp' },
    { name: 'AUDI', logoPath: '/Audi-Logo-Banner.avif' },
    { name: 'MERCEDES', logoPath: '/Mercedes-Logo-.svg.png' }
]

export default function Cars() {
    const [addCar, setAddCar] = useState(false)
    function toggleModal() {
        setAddCar(!addCar)
    }

    return (
        <div className={styles.carsContainer}>
            {
                !addCar &&
                <div onClick={toggleModal} className={styles.addCarBtn}>
                    <FontAwesomeIcon icon={faPlus} />
                    {'  '}Add New Car
                </div>
            }
            <div className={styles.addModal}>
                {addCar && <AddCarModal toggleModal={toggleModal} />}
            </div>
            <h1 style={{ margin: '10px' }}>
                Cars Management
            </h1>
            <div>
                <label className={styles.label}>
                    Status:{' '}
                    <select className={styles.select}>
                        <option>All</option>
                        <option>Available</option>
                        <option>Rented</option>
                    </select>
                    <FontAwesomeIcon className={styles.down} icon={faCaretDown} />
                </label>
                <label className={styles.label}>
                    Brands:{' '}
                    <select className={styles.select}>
                        <option>All</option>
                        {
                            brands.map((el, i) => <option key={i}>{el.name}</option>)
                        }
                    </select>
                    <FontAwesomeIcon className={styles.down} icon={faCaretDown} />
                </label>
                <label className={styles.label}>
                    Types:{' '}
                    <select className={styles.select}>
                        <option>All</option>
                        {
                            types.map((el, i) => <option key={i}>{el}</option>)
                        }
                    </select>
                    <FontAwesomeIcon className={styles.down} icon={faCaretDown} />
                </label>
            </div>
        </div>
    )
}