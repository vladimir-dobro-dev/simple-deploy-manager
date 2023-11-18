import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'
import LanguageDetector from 'i18next-browser-languagedetector';

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    debug: true,
    // lng: "ru",
    fallbackLng: 'en',
    interpolation: {
      escapeValue: false,
    },
    resources: {
      en: {
        Installer: {
          Path: "Installation path",
          Install: "Install",
          Exit: "Exit",
        },
      },
      ru: {
        Installer: {
            Path: "Путь для установки",
            Install: "Установить",
            Exit: "Выход",
        },
      },
    }
  });

export default i18n;
