import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCar, faTag, faBuilding, faAddressBook, faShop } from '@fortawesome/free-solid-svg-icons';
import { useState } from 'react';
import styles from "./sideBar.module.css"

export default function SidBar() {
    const [navItem, setNavItem] = useState('cars');

    return (
        <div className={styles.sideBar}>
            <div className={styles.header}>Car Rental Admin</div>
            <div
                className={navItem === "cars" ? `${styles.sideBarItem} ${styles.selected}` : styles.sideBarItem}
                onClick={() => setNavItem('cars')}
            >
                <FontAwesomeIcon className={styles.icon} icon={faCar} />
                <span className={styles.tag}>Cars</span>
            </div>
            <div
                className={navItem === "types" ? `${styles.sideBarItem} ${styles.selected}` : styles.sideBarItem}
                onClick={() => setNavItem('types')}
            >
                <FontAwesomeIcon className={styles.icon} icon={faTag} />
                <span className={styles.tag}>Types</span>
            </div>
            <div
                className={navItem === "brands" ? `${styles.sideBarItem} ${styles.selected}` : styles.sideBarItem}
                onClick={() => setNavItem('brads')}
            >
                <FontAwesomeIcon className={styles.icon} icon={faBuilding} />
                <span className={styles.tag}>Brands</span>
            </div>
            <div
                className={navItem === "contacts" ? `${styles.sideBarItem} ${styles.selected}` : styles.sideBarItem}
                onClick={() => setNavItem('contacts')}
            >
                <FontAwesomeIcon className={styles.icon} icon={faAddressBook} />
                <span className={styles.tag}>Contacts</span>
            </div>
            <div
                className={navItem === "info" ? `${styles.sideBarItem} ${styles.selected}` : styles.sideBarItem}
                onClick={() => setNavItem('info')}
            >
                <FontAwesomeIcon className={styles.icon} icon={faShop} />
                <span className={styles.tag}>Shop Info</span>
            </div>
        </div>
    )

}