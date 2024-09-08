import config from '~/configs';
import defaultLayout from '~/layout/defaultLayout';
import OnlyChildren from '~/layout/onlyChildren';
import Home from '~/pages/home/home';
import Login from '~/pages/login';
import Register from '~/pages/register';
const publicRoutes = [
    {
        path: config.routes.home,
        component: Home,
        layout: defaultLayout,
    },
    {
        path: config.routes.login,
        component: Login,
        layout: OnlyChildren,
    },
    {
        path: config.routes.register,
        component: Register,
        layout: OnlyChildren,
    },
];

export { publicRoutes };
