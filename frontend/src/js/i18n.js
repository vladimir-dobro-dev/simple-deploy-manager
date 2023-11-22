import i18next from "i18next"
import LanguageDetector from "i18next-browser-languagedetector"

i18next
    .use(LanguageDetector)
    .init({
        debug: true,
        lng: "ru",
        fallbackLng: 'en',
        interpolation: {
            escapeValue: false,
        },
        resources: {
            en: {
                translation: {
                    ServerConnect: {
                        Title: "Add server",
                        IPAddress: "IP address",
                        Name: "User name",
                        Password: "Password",
                        Port: "Port",
                        ButtonConnect: "Connect to server",
                    },
                }
            },
            ru: {
                translation: {
                    ServerConnect: {
                        Title: "Добавление сервера",
                        IPAddress: "IP адрес",
                        Name: "Имя пользователя",
                        Password: "Пароль",
                        Port: "Порт",
                        ButtonConnect: "Подключиться к серверу",
                    },
                }
            },
        }
    });

export default i18next
