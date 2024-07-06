import React from "react";
import "../pages/Popup.css"; // Reuse the Popup styles

interface LogoutConfirmationPopupProps {
    username: string | null;
    onConfirm: () => void;
    onCancel: () => void;
}

const LogoutConfirmationPopup: React.FC<LogoutConfirmationPopupProps> = ({
    username,
    onConfirm,
    onCancel,
}) => {
    return (
        <div className="popup">
            <div className="popup-content">
                <h3>{username}, are you sure you want to sign out?</h3>
                <div className="popup-buttons">
                    <button className="button-style" onClick={onConfirm}>
                        Yes
                    </button>
                    <button className="button-style" onClick={onCancel}>
                        No
                    </button>
                </div>
            </div>
        </div>
    );
};

export default LogoutConfirmationPopup;
