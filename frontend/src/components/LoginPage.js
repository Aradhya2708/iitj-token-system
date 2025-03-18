import React, { useState } from 'react';
import './LoginPage.css';
import Navbar from './Navbar';
import { useNavigate } from 'react-router-dom'; 
import axios from 'axios';

function LoginPage() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();

        try {
            const response = await axios.post('http://localhost:8080/login', {
                username,
                password,
            });

            if (response.status === 200) {
                console.log('User authenticated:', response.data);
                navigate('/student');
            }
        } catch (error) {
            console.error('Login failed:', error);
            if (error.response && error.response.data) {
                setErrorMessage(error.response.data.message || 'Invalid credentials');
            } else {
                setErrorMessage('Failed to connect to the server');
            }
        }
    };

    return (
        <>
            <Navbar links={[]} />
            <div className="card text-center Psize container">
                <div className="card-header">IIT-J Token System</div>
                <div className="card-body">
                    <h5 className="card-title">Enter your LDAP credentials to proceed</h5>
                    {errorMessage && <p className="error-message">{errorMessage}</p>}
                    <form onSubmit={handleLogin}>
                        <div className="userName">
                            <input
                                type="text"
                                placeholder="Username"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                                required
                            />
                        </div>
                        <div className="password">
                            <input
                                type="password"
                                placeholder="Password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                            />
                        </div>
                        <button type="submit" className="btn btn-info">
                            Log In
                        </button>
                    </form>
                </div>
                <div className="card-footer text-muted">~ DevlUp Labs</div>
            </div>
        </>
    );
}

export default LoginPage;
