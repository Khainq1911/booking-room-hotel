import { checkLogin } from '~/services/loginService';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser } from '@fortawesome/free-regular-svg-icons';
import { faLock } from '@fortawesome/free-solid-svg-icons';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
function Login() {
    const [userName, setUserName] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(false);

    const [checkPhone, setCheckPhone] = useState(false);
    const [checkPassword, setCheckPassword] = useState(false);

    const navigate = useNavigate();
    const data = {
        phone: userName,
        password: password,
    };

    // call api to login
    const login = async () => {
        // check phone and password
        if (data.phone.trim() === '') {
            setCheckPhone(true);
        }
        if (data.password.trim() === '') {
            setCheckPassword(true);
        }

        try {
            //call api
            const res = await checkLogin(JSON.stringify(data));
            localStorage.setItem('token', res.data.token);
            navigate('/');

            return res;
        } catch (error) {
            setError(true);
        }
    };

    return (
        <div className="bg-[#99CCFF] w-full h-screen flex justify-center items-center">
            <div className="bg-[white] w-[300px] p-6 rounded-lg shadow-lg">
                <h1 className="text-2xl text-center mb-2">Login</h1>
                {error && <p className="text-[14px] text-red-500 mb-2">Phone or password is false. Try again</p>}
                <div>
                    <div className="mb-4 relative">
                        <label className="block text-[14px]" htmlFor="username">
                            Phone
                        </label>
                        <input
                            type="text"
                            required
                            placeholder="type your phone number"
                            value={userName}
                            onChange={(e) => {
                                setUserName(e.target.value);
                            }}
                            onFocus={() => (setError(false), setCheckPhone(false))}
                            className="w-full rounded-[10px] border-b pl-[25px] pr-[8px] py-[4px] focus:outline-gray-400 border-gray-300 mt-[4px]"
                        />
                        {checkPhone && <p className="text-[14px] text-red-500 mb-2">Phone number is empty</p>}
                        <FontAwesomeIcon
                            icon={faUser}
                            className="size-3 text-gray-500 absolute top-[33px] left-[4px]"
                        />
                    </div>
                    <div className="mb-4 relative">
                        <label className="block text-[14px]" htmlFor="password">
                            Password
                        </label>
                        <input
                            type="password"
                            required
                            onChange={(e) => {
                                setPassword(e.target.value);
                            }}
                            onFocus={() => (setError(false), setCheckPassword(false))}
                            placeholder="type your password"
                            className="w-full rounded-[10px] border-b pl-[25px] pr-[8px] py-[4px] focus:outline-gray-400 border-gray-300 mt-[4px]"
                        />
                        <FontAwesomeIcon
                            icon={faLock}
                            className="size-3 text-gray-500 absolute top-[33px] left-[4px]"
                        />
                        {checkPassword && <p className="text-[14px] text-red-500 mb-2">Password number is empty</p>}
                        <a href="/" className="block text-right text-[12px] mt-[8px]">
                            Forgot your password?
                        </a>
                    </div>
                    <button
                        type="submit"
                        className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 w-full"
                        onClick={login}
                    >
                        Login
                    </button>
                    <a href="/register" className="block mt-[20px] text-center">
                        Create account?
                    </a>
                </div>
            </div>
        </div>
    );
}

export default Login;
