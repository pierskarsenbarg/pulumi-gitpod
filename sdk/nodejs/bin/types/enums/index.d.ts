export declare const EditorName: {
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Code: "code";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Code_desktop: "code-desktop";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Intellij: "intellij";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Goland: "goland";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Phpstorm: "phpstorm";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Pycharm: "pycharm";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Rubymine: "rubymine";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Webstorm: "webstorm";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Rider: "rider";
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Clion: "clion";
};
export type EditorName = (typeof EditorName)[keyof typeof EditorName];
export declare const EditorVersion: {
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly Latest: "latest";
    /**
     * Standard workspace class - up to 4 cores, up to 8GB RAM, 30GB storage
     */
    readonly Stable: "stable";
};
export type EditorVersion = (typeof EditorVersion)[keyof typeof EditorVersion];
export declare const WorkspaceClass: {
    /**
     * Large workspace class - up to 8 cores, up to 16GB RAM, 50GB storage.
     */
    readonly G1_large: "g1-large";
    /**
     * Standard workspace class - up to 4 cores, up to 8GB RAM, 30GB storage
     */
    readonly G1_standard: "g1-standard";
};
export type WorkspaceClass = (typeof WorkspaceClass)[keyof typeof WorkspaceClass];
