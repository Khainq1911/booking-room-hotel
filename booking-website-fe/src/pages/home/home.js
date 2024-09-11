import { useNavigate } from 'react-router-dom';
import { logout } from '~/services/loginService';

function Home() {
    const navigate = useNavigate();
    const logoutAction = async () => {
        try {
            const res = await logout();
            localStorage.removeItem('token');
            navigate('/login');
            return res;
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <div>
            <button onClick={logoutAction} className="text-[20px] bg-[red]">
                logout
            </button>
        </div>
    );
}
export default Home;
