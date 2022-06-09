import Home from "@/components/Home";
import Login from "@/components/Login";
import VueRouter from "vue-router";
import ClusterList from "../components/cluster/ClusterList";
import InstanceList from "../components/cluster/InstanceList";
import TaskList from "../components/task/TaskList";
import DbInstanceList from "../components/task/DbInstanceList";
import DataSourceList from "../components/datasource/DataSourceList";
import DataSourceDetail from "../components/datasource/DataSourceDetail";

const routes = [
    {
        path: "/",
        name: "home",
        component: Home,
        children: [{
            path: "/cluster/all",
            name: "cluster_list",
            component: ClusterList
        }, {
            path: "/cluster/all",
            name: "instance_list",
            component: InstanceList
        }, {
            path: "/worker/all",
            name: "task_list",
            component: TaskList
        }, {
            path: "/worker/all",
            name: "db_instance_list",
            component: DbInstanceList
        }, {
            path: "/datasource/all",
            name: "db_datasource_list",
            component: DataSourceList
        },{
            path:"/datasource/detail",
            name:"db_datasource_detail",
            component: DataSourceDetail
        }]
    },
    {
        path: "/login",
        name: "login",
        component: Login
    }
]

export const router = new VueRouter({
    routes
})
