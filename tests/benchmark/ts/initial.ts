import * as api from "./api"

const teacherNum = 10;
const studentNum = 50;
const termNum = 3;
const classNameNum = 5;
const recNum = 10;
const admin = new api.User("admin", "123456")

let teachers = new Array<api.User>(teacherNum);

const adminInit = async () => {
    let adminToken: string;
    adminToken = await api.captcha()
        .then(x => api.login(admin.account, admin.password, x.captchaId, x.d))
    for (let index = 0; index < teacherNum; index++) {
        teachers[index] = await api.newTercher(adminToken)
    }
    let begin = 2020;
    for (let index = 0; index < termNum; index++, begin++) {
        api.newTerm(adminToken, `${begin}第一学期`, `${begin}-03-01`, `${begin}-07-01`)
    }
    for (let index = 0; index < classNameNum; index++) {
        api.newCN(adminToken, api.randomString(8))
    }
};

const student = async (account: string, password: string, eid: number) => {
    let token: string;
    token = await api.captcha()
        .then(x => api.login(account, password, x.captchaId, x.d))
    return Promise.all([api.postFile("test.zip"), api.postImg("cat.jpg")])
        .then(([res1, res2]) => api.submit(token, res1 as string, res2 as string, eid))
}

const teacherInit = async (i: number) => {
    let teacherToken: string;
    teacherToken = await api.captcha()
        .then(x => api.login(teachers[i].account, teachers[i].password, x.captchaId, x.d))

    for (let index = 1; index <= termNum; index++) {
        for (let j = 1; j <= classNameNum; j++) {
            let students = new Array<string>(studentNum);
            for (let k = 0; k < studentNum; k++) {
                students[k] = api.randomString(8)
            }
            await api.newClass(teacherToken, j, index)
                .then(cid => {
                    api.addStudents(teacherToken, cid, students)
                        .then(() =>
                            api.newExp(teacherToken, cid)
                                .then(async eid => {
                                    for (let k = 0; k < studentNum; k++) {
                                        await student(students[k], "666666", eid)
                                            .then(() => { if (k < recNum) api.rec(teacherToken, students[k], eid); })
                                    }
                                })
                        )
                })

        }
    }
};

let main = async () => {
    await adminInit()
    console.log(teachers)
    for (let index = 0; index < teacherNum; index++) {
        await teacherInit(index)
    }
};

main();

