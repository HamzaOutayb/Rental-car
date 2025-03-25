'use client'
import { useState } from 'react';
import AddCarModal from '../components/addCarModal';
import SidBar from './app/components/sideBar';
import  Cars from '../components/cars'

export default function Dashboard() {
  const [isModalOpen, setModalOpen] = useState(false);
  var cars = 'Cars'
  const toggleModal = () => {
    setModalOpen(!isModalOpen);
  };

  return (
    <div style={{backgroundColor:'#D0DEE9', width:'100vw', height:'100vh', margin:'0', padding:'0'}}>
      <SidBar></SidBar>
      <Cars/>
    </div>
    // <div className="dashboard">
    //   <h1>Admin Dashboard</h1>
    //   <button className="add-car-btn" onClick={toggleModal}>
    //     Add Car
    //   </button>

    //   {/* Modal for adding a car */}
    //   {isModalOpen && <AddCarModal toggleModal={toggleModal} />}

    //   {/* Car List */}
    //   <div className="car-list">
    //     <h2>Existing Cars</h2>
    //     <table>
    //       <thead>
    //         <tr>
    //           <th>Car Name</th>
    //           <th>Status</th>
    //           <th>Actions</th>
    //         </tr>
    //       </thead>
    //       <tbody>
    //         {/* Add your car data rows here */}
    //         <tr>
    //           <td>Car 1</td>
    //           <td>Available</td>
    //           <td><button>Edit</button></td>
    //         </tr>
    //         <tr>
    //           <td>Car 2</td>
    //           <td>Rented</td>
    //           <td><button>Edit</button></td>
    //         </tr>
    //       </tbody>
    //     </table>
    //   </div>
    // </div>
  );
}